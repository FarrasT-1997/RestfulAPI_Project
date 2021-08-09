echo "Run and create container (image should be exist)"
docker run -d --name altaContainer -p 3000:8080 alta-ecommerce:latest