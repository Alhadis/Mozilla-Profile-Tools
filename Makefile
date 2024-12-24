all: mozinstallhash mozlz4

# Program for computing hashes used in Mozilla installation profiles
mozinstallhash: mozinstallhash.go
	go build $^
	./$@ Foo | grep -qiF 9A12EB455E563003


# MozLZ4-encoder and decoder tool
mozlz4 = mozlz4-src/target/release/mozlz4
mozlz4: $(mozlz4)
	mv $^ $@
	./$@ --help | grep -iqF 'Decompress and compress mozlz4 files'

mozlz4-src:
	git clone https://github.com/jusw85/mozlz4.git $@

$(mozlz4): mozlz4-src
	cd $^ && cargo build --release
