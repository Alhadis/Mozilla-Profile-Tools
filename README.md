Mozilla profile tools
================================================================================
This repository contains build instructions for two *extremely* niche utilities:

* [`mozinstallhash`]: A program to compute the hash used in Mozilla installation
  profile names. These form part of the unique identifier strings that represent
  a user profile directory (e.g., in `~/Application Support/Firefox/Profiles/`).
  [^1]‍[^2]

* [`mozlz4`]: A program to compress and decompress files using Mozilla's variant
  of the [LZ4 algorithm](https://w.wiki/AeWU). This is used by (recent) releases
  of Firefox to compress certain JSON files in user profile directories.

Both of these tools reference poorly-documented facilities in Mozilla's codebase
that are of interest to those who programmatically manipulate their user profile
directories outside of Firefox.


Usage
--------------------------------------------------------------------------------
Building these tools requires [Go] (v1.15 or newer) and [Rust] (v1.61 or newer).
To build and install them, run:

~~~console
$ make all
$ make install                          # Default install locations
$ make bindir=/path/to/dir install      # Specify location of installed binaries
~~~

The `make install` step is optional and requires a BSD-style [`install(1)`][] to
be available on the user's system. Users can "install" the binaries simply using
[`mv(1)`][]; no other files are required to run the programs once they're built.


Mirrors
--------------------------------------------------------------------------------
In the event that an upstream repository "disappears", a copy of its source code
(possibly outdated) can be found in a detached branch of this repository. To use
these backups, modifications to the source code may be required (details of such
changes are beyond the scope of this document).


<!-- Footnotes ---------------------------------------------------------------->
[^1]: https://github.com/twpayne/chezmoi/issues/1226#issuecomment-867228095 \
      A well-researched overview of Mozilla's use of hashing to generate profile
      ID strings. Courtesy of [**@bradenhilton**], who also authored the Go code
      used by this repository (with minor modifications).

[^2]: https://bit.ly/41NXOC7 \
      Shortlink to `toolkit/mozapps/update/common/commonupdatedir.cpp` (revision
      <samp>2fcb225</samp>) in the online code browser for the `mozilla-central`
      repository.

<!-- Referenced links --------------------------------------------------------->
[**@bradenhilton**]: https://github.com/bradenhilton
[`mozinstallhash`]:  https://github.com/bradenhilton/mozillainstallhash
[`mozlz4`]:          https://github.com/jusw85/mozlz4
[`mv(1)`]:           https://man.openbsd.org/mv.1
[`install(1)`]:      https://man.openbsd.org/install.1
[Rust]:              https://www.rust-lang.org/tools/install
[Go]:                https://go.dev/dl/

<!-- Editor settings -----------------------------------------------------------
Local Variables:
  coding: utf-8-unix
  mode: GFM
  tab-width: 2
  fill-column: 80
  indent-tabs-mode: nil
  truncate-lines: t
End:
vim: ff=unix ft=gfm ts=2 tw=80 et nowrap
------------------------------------------------------------------------------->
