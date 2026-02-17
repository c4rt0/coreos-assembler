---
parent: CoreOS Assembler Command Line Reference
nav_order: 3
---

# cosa diff

`cosa diff` compares two builds in a cosa workdir. By default it compares the
previous build to the latest build. You can specify explicit builds with
`--from` and `--to`, and select an architecture with `--arch`

```
$ cosa diff --from 42.20260131.dev.0 --to 42.20260201.dev.0
```

At least one differ flag must be passed to tell the script what needs to be compared.

## Available differs

  `-h, --help`            show this help message and exit
  `--from DIFF_FROM`      First build ID
  `--to DIFF_TO`          Second build ID
  `--gc`                  Delete cached diff content
  `--arch ARCH`           Architecture of builds
  `--difftool`            Use git difftool
  `--rpms-rpm-ostree`     Diff RPMs using rpm-ostree
  `--rpms-rpm-ostree-json`
                          Diff RPMs & Advisories using rpm-ostree, output JSON
  `--rpms`                Diff rpms from commitmeta.json
  `--rpms-json`           Diff RPMS & Advisories from commitmeta.json, output JSON
  `--source-control`      Diff config and COSA input commits
  `--ostree-ls`           Diff OSTree contents using 'ostree diff'
  `--ostree`              Diff OSTree contents using 'git diff'
  `--initrd`              Diff initramfs contents
  `--live-iso-ls`         Diff live ISO listings
  `--live-iso`            Diff live ISO content
  `--live-initrd-ls`      Diff live initramfs listings
  `--live-initrd`         Diff live initramfs content
  `--live-rootfs-img-ls`  Diff live-rootfs.img listings
  `--live-rootfs-img`     Diff live-rootfs.img content
  `--live-sysroot-ls`     Diff live '/root.[erofs|squash]fs' (embed into live-rootfs) listings
  `--live-sysroot`        Diff live '/root.[ero|squash]fs' (embed into live-rootfs) content
  `--metal-part-table`    Diff metal disk image partition tables
  `--metal`               Diff metal disk image content
  `--metal-du`            Compare directory usage of metal disk image content
  `--metal-ls`            Compare directory listing of metal disk image content
  `--sizes`               Diff artifact sizes from meta.json

## Combining differs

Multiple differs can be passed in at once:

```
$ cosa diff --rpms --sizes
```

## Examples

Compare RPMs between the two most recent builds:

```
$ cosa diff --rpms
```

Compare artifact sizes for specific builds:

```
$ cosa diff --sizes --from 42.20260131.91.0 --to 42.20260201.dev.0
```

Compare OSTree content between latest and a specific build:

```
$ cosa diff --ostree --from 42.20250519.dev.0
```
