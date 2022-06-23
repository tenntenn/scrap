# trytoolexec

https://zenn.dev/tenntenn/scraps/bb07822b180526

## go run

```
$ go run -toolexec="go run ./cmd/mytool" .
# github.com/tenntenn/scrap/trytoolexec/b
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/b
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b003/_pkg_.a -trimpath $WORK/b003=> -p github.com/tenntenn/scrap/trytoolexec/b -lang=go1.18 -complete -buildid E1q9GG9Uekl0mQ8YOUEL/E1q9GG9Uekl0mQ8YOUEL -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b003/importcfg -pack b/b.go
# github.com/tenntenn/scrap/trytoolexec/a
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/a
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b002/_pkg_.a -trimpath $WORK/b002=> -p github.com/tenntenn/scrap/trytoolexec/a -lang=go1.18 -complete -buildid g_xDdkA9h2Bky2Yh_qCT/g_xDdkA9h2Bky2Yh_qCT -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b002/importcfg -pack a/a.go
# github.com/tenntenn/scrap/trytoolexec
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b001/_pkg_.a -trimpath $WORK/b001=> -p main -lang=go1.18 -complete -buildid PL3ZMCZ0RPoX2ddfp-4M/PL3ZMCZ0RPoX2ddfp-4M -dwarf=false -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b001/importcfg -pack ./main.go
# github.com/tenntenn/scrap/trytoolexec
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/b001/exe/trytoolexec -importcfg $WORK/b001/importcfg.link -s -w -buildmode=exe -buildid=TagZU7ogkDn5Ycd6TPHO/PL3ZMCZ0RPoX2ddfp-4M/tW26ZGQl5pEMYZWjzQTZ/TagZU7ogkDn5Ycd6TPHO -extld=clang $WORK/b001/_pkg_.a
done
```

## go test

```
$ go test -vet off -count 1 -v -toolexec="go run ./cmd/mytool" ./...
# github.com/tenntenn/scrap/trytoolexec/a
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/a
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b002/_pkg_.a -trimpath $WORK/b002=> -p github.com/tenntenn/scrap/trytoolexec/a -lang=go1.18 -complete -buildid g_xDdkA9h2Bky2Yh_qCT/g_xDdkA9h2Bky2Yh_qCT -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b002/importcfg -pack a/a.go
# github.com/tenntenn/scrap/trytoolexec/b
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/b
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b003/_pkg_.a -trimpath $WORK/b003=> -p github.com/tenntenn/scrap/trytoolexec/b -lang=go1.18 -complete -buildid E1q9GG9Uekl0mQ8YOUEL/E1q9GG9Uekl0mQ8YOUEL -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b003/importcfg -pack b/b.go
# github.com/tenntenn/scrap/trytoolexec/cmd/mytool
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/cmd/mytool
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b085/_pkg_.a -trimpath $WORK/b085=> -p main -lang=go1.18 -complete -buildid BBHnDMSiBO81veA2XVdb/BBHnDMSiBO81veA2XVdb -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b085/importcfg -pack cmd/mytool/main.go
# github.com/tenntenn/scrap/trytoolexec/a [github.com/tenntenn/scrap/trytoolexec/a.test]
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/a [github.com/tenntenn/scrap/trytoolexec/a.test]
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b015/_pkg_.a -trimpath $WORK/b015=> -p github.com/tenntenn/scrap/trytoolexec/a -lang=go1.18 -complete -buildid 0bEdiO_awwdY0Ifzf9FK/0bEdiO_awwdY0Ifzf9FK -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b015/importcfg -pack a/a.go a/a_test.go
# github.com/tenntenn/scrap/trytoolexec
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b001/_pkg_.a -trimpath $WORK/b001=> -p main -lang=go1.18 -complete -buildid bmBIOPviORRKdRWybgE6/bmBIOPviORRKdRWybgE6 -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b001/importcfg -pack ./main.go
?   	github.com/tenntenn/scrap/trytoolexec	[no test files]
# github.com/tenntenn/scrap/trytoolexec/b [github.com/tenntenn/scrap/trytoolexec/b.test]
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/b [github.com/tenntenn/scrap/trytoolexec/b.test]
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b082/_pkg_.a -trimpath $WORK/b082=> -p github.com/tenntenn/scrap/trytoolexec/b -lang=go1.18 -complete -buildid Fo7OVNprdaEmIpKL3udf/Fo7OVNprdaEmIpKL3udf -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b082/importcfg -pack b/b.go b/b_test.go
# github.com/tenntenn/scrap/trytoolexec/a.test
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/a.test
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b053/_pkg_.a -trimpath $WORK/b053=> -p main -lang=go1.18 -complete -buildid zUIvuQdiJdOYtsXkK0kn/zUIvuQdiJdOYtsXkK0kn -dwarf=false -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b053/importcfg -pack $WORK/b053/_testmain.go
# github.com/tenntenn/scrap/trytoolexec/b.test
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/b.test
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b083/_pkg_.a -trimpath $WORK/b083=> -p main -lang=go1.18 -complete -buildid 5YvcnMtLtPQLhVUl59Uz/5YvcnMtLtPQLhVUl59Uz -dwarf=false -goversion go1.18 -c=4 -nolocalimports -importcfg $WORK/b083/importcfg -pack $WORK/b083/_testmain.go
# github.com/tenntenn/scrap/trytoolexec/a.test
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/a.test
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/b053/a.test -importcfg $WORK/b053/importcfg.link -s -w -buildmode=exe -buildid=whu4R_hH6ihkSbx8t-zj/zUIvuQdiJdOYtsXkK0kn/SjN_2cOVM8fj5bKNMvC2/whu4R_hH6ihkSbx8t-zj -extld=clang $WORK/b053/_pkg_.a
# github.com/tenntenn/scrap/trytoolexec/b.test
TOOLEXEC_IMPORTPATH github.com/tenntenn/scrap/trytoolexec/b.test
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/b083/b.test -importcfg $WORK/b083/importcfg.link -s -w -buildmode=exe -buildid=lAwc1uANkCEk812RgxOF/5YvcnMtLtPQLhVUl59Uz/bC-wuKNoza0aUlcGU4v1/lAwc1uANkCEk812RgxOF -extld=clang $WORK/b083/_pkg_.a
=== RUN   TestA
--- PASS: TestA (0.00s)
PASS
ok  	github.com/tenntenn/scrap/trytoolexec/a	0.931s
=== RUN   TestB
--- PASS: TestB (0.00s)
PASS
ok  	github.com/tenntenn/scrap/trytoolexec/b	1.516s
?   	github.com/tenntenn/scrap/trytoolexec/cmd/mytool	[no test files]
```

## -exec

-execはgo runやgo testで裏でビルドしたバイナリを実行する際に変わりに別のバイナリを実行するオプション

```
$ go run -exec echo .
/var/folders/dr/gn16sh9935zcz37xt2ghpk480000gn/T/go-build3356048058/b001/exe/trytoolexec
```

```
$ go test -vet off -count 1 -v -exec echo ./...
?   	github.com/tenntenn/scrap/trytoolexec	[no test files]
/var/folders/dr/gn16sh9935zcz37xt2ghpk480000gn/T/go-build3150231640/b053/a.test -test.paniconexit0 -test.timeout=10m0s -test.count=1 -test.v=true
ok  	github.com/tenntenn/scrap/trytoolexec/a	0.005s
/var/folders/dr/gn16sh9935zcz37xt2ghpk480000gn/T/go-build3150231640/b083/b.test -test.paniconexit0 -test.timeout=10m0s -test.count=1 -test.v=true
ok  	github.com/tenntenn/scrap/trytoolexec/b	0.004s
?   	github.com/tenntenn/scrap/trytoolexec/cmd/mytool	[no test files]
```

## delete cache

```
$ go clean -cache
```
