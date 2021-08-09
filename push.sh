# ini bergantung pada username dockerhub
USERNAME=anggunpermatata

docker build -t $USERNAME/alta-ecommerce:latest .
docker push $USERNAME/alta-ecommerce:latest

docker run -d --name altaContainer -p 3000:8080 $USERNAME/alta-ecommerce:latest