targets = mozinstallhash mozlz4

all: $(targets)


# Program for computing hashes used in Mozilla installation profiles
mozinstallhash: mozinstallhash.go
	go build $^
	./$@ Foo | grep -qiF 9A12EB455E563003


# MozLZ4-encoder and decoder tool
mozlz4 = mozlz4-src/target/release/mozlz4
mozlz4: $(mozlz4)
	ln -f $^ $@
	./$@ --help | grep -iqF 'Decompress and compress mozlz4 files'

mozlz4-src:
	git clone https://github.com/jusw85/mozlz4.git $@

$(mozlz4): mozlz4-src
	cd $^ && cargo build --release


# Use a BSD-style install(1) to move the built programs to one's $PATH
prefix = /usr/local
bindir = ${prefix}/bin

install: $(targets)
	install -vpC $^ "${bindir}"

.PHONY: install


# Purge directory of generated and untracked files
clean:
	rm -rf mozlz4-src
	rm -f ${targets}
	if test -d sources; then \
		chmod -R u+w sources; \
		rm -rf sources; \
	fi

.PHONY: clean
