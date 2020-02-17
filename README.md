# open-pull-request-URL
Can open Pull Request(PR) creation form easily after pusing your branch.

## How To Use
You can open PR creation form after git push
```
git checkout -b test-branch
echo "Hello" > test.txt
git add test.txt
git commit -m "test"
git push origin HEAD
pr
```

## How To Setup
This CLI is created with golang.
Please check your GOPATH and PATH before setup.
```
cd $GOPATH

git clone https://github.com/kappyhappy/open-pull-request-URL.git

cd open-pull-request-URLopen-pull-request-URLopen-pull-request-URLopen-pull-request-URL

go build .

mv pr $PATH
```
