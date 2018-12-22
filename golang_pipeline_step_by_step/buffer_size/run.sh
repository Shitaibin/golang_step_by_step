n=0
while(($n<10))
do
    name="BenchmarkPipelineFanBuffered_"$n
    echo $name
    go test -test.bench=$name
    echo ""
    let "n++"
done