targets = mozinstallhash mozlz4

all: $(targets)


# Program for computing hashes used in Mozilla installation profiles
mozinstallhash: mozinstallhash.go
	go build $?
	./$@ Foo | grep -qiF 9A12EB455E563003


# MozLZ4-encoder and decoder tool
mozlz4: mozlz4-src
	cd $? && cargo build --release
	mv -v mozlz4-src/target/release/mozlz4 $@
	./$@ --help | grep -iqF 'Decompress and compress mozlz4 files'
	cd $? && cargo clean -vv && ${resetRepoTimestamp}

mozlz4-src: mozlz4-src/.git
mozlz4-src/.git:
	git submodule init
	git submodule sync --quiet
	git submodule update -- $(@D)

# Set the access and modification times of a submodule to its last commit
resetRepoTimestamp = \
	git show --no-patch --format='%at %ct' \
	| xargs perl -E 'my($$t)=sort{$$b<=>$$a}(shift,shift);utime $$t,$$t,"."';


# Use a BSD-style install(1) to move the built programs to one's $PATH
prefix = /usr/local
bindir = ${prefix}/bin

install: $(targets)
	install -vpC ${targets} "${bindir}"

.PHONY: install


# Purge directory of generated and untracked files
clean:
	rm -f ${targets}
	if test -d sources; then \
		chmod -R u+w sources; \
		rm -rf sources; \
	fi

.PHONY: clean


# Verify expected behaviour and output of compiled utilities
test: \
	test-mozinstallhash \
	test-mozlz4

test-mozinstallhash: mozinstallhash
	@ ${testHeader} \
	expectHash(){ \
		set -- "$$@" "`./$? "$$2"`"; \
		case "$$3" in "$$1");; *) \
			printf >&2 'Assertion failed for "%s":\n' "$$2"; \
			printf >&2 '\tExpected: %s\n' "$$1"; \
			printf >&2 '\tActual:   %s\n' "$$3"; \
			return 1 ;; \
		esac; \
	}; \
	expectHash 308046B0AF4A39CB "C:\Program Files\Mozilla Firefox"; \
	expectHash 9D561FCD08DC6D55 "C:/Program Files/Mozilla Firefox"; \
	expectHash 4F96D1932A9F858E '/usr/lib/firefox'; \
	expectHash 2656FF1E876E9973 '/Applications/Firefox.app/Contents/MacOS';
	@ ${testFooter}

test-mozlz4: mozlz4
	@ ${testHeader} \
	rm -rf fixtures/*.baklz4; \
	set -e; \
	set -- 'fixtures/search.json'; \
	./$? --compress "$$1" "$$1.mozlz4"; \
	test "`size "$$1"`"        -eq 529; \
	test "`size "$$1.mozlz4"`" -eq 327; \
	cd fixtures && $$sha256c SHA256; \
	set -- "$${1#*/}.mozlz4"; \
	../$? --extract "$$1" "$$1.baklz4"; \
	test "`size "$$1.baklz4"`" -eq "`size "$${1%.*}"`"; \
	hashesMatch "$$1.baklz4" "$${1%.*}"; \
	../$? --compress search.json -        | startsWith 'mozLz40'; \
	../$? --extract  search.json.mozlz4 - | startsWith '{"version":10,'; \
	../$? --compress - tmp.baklz4 < search.json; \
	../$? --extract  - tmp.json   < tmp.baklz4; \
	hashesMatch tmp.baklz4 search.json.mozlz4; \
	hashesMatch tmp.json   search.json;
	@ ${testFooter}

# Helper functions shared by test-* tasks
# TODO: Use a more elegant workaround for `--quiet` switches
#       passed to sha256sum(1) when not verifying checksums.
testHeader = \
	echo "Testing $?..."; \
	if command -v sha256sum >/dev/null 2>&1; \
		then sha256='sha256sum'; sha256c="$$sha256 --quiet -c"; \
		else sha256='sha256 -r'; sha256c="$$sha256 -qc"; \
	fi; \
	case `stat --version 2>/dev/null` in \
		*GNU*) size(){ stat -c %s "$$1"; } ;; \
		*)     size(){ stat -f %z "$$1"; } ;; \
	esac; \
	hashesMatch(){ \
		count=`$$sha256 "$$@" | cut -d' ' -f1 | uniq | wc -l`; \
		count=`printf %s "$$count" | expand | sed 's/^ *//; s/ *$$//;'`; \
		test "$$count" -eq 1; \
	}; \
	startsWith(){ \
		set -- "$$1" "$${\#1}"; \
		test "$$2" -gt 0; \
		dd 2>/dev/null bs="$$2" count=1 | grep -qF "$$1"; \
	};

# Commands executed after a completed test-* task
testFooter = \
	echo "Tests passed for $?!"; \
	git clean -xfd fixtures
