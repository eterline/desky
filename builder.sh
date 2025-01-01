#!/bin/sh

mkdir logging

echo "Building backend app"
cd desky-backend/
go build -o app -v ./cmd/desky-backend/main.go

echo "Copy backend app files and dirs"
cp app ../
cp apps.json ../


echo "Generating backend configuration file"
cd ..
./app -generate true
rm trace.log


cd desky-front/

echo "Installing NPM dependencies"
npm install
echo "Building frontend static files"
npm run build:prod

echo "Copy frontend static files to main directory"
cp -R web ../

cd ..
echo "Create starting file"
touch run.sh

tee -a run.sh <<EOF
#!/bin/sh

./app -log logging
EOF
chmod +x run.sh

echo "Done."
