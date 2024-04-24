echo "building UI" && \
cd ui && \
npm run build && \
cd .. && \
rm -rf build && \
mkdir build && \
cp ui/build/* build -r && \
echo "building server" && \
cd src && \
go build && \
mv src main
