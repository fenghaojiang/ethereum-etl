
# find all abi file and generate go resolve file in ./onchain/generated-go directory
for fileName in $(find ./onchain/abis/ -name "*.abi" -type f); do
    dir="./onchain/generated-go/$(dirname "${fileName#./onchain/abis/}")"
    mkdir -p "$dir"
    bname="$(basename "$fileName" .abi)"
    abigen --abi="$fileName" --pkg="$bname" --out="$dir/$bname.go"
done
