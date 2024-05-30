docker pull postgres:16.3

docker run --name deli-blog-postgres-container -p 5432:5432 -e POSTGRES_PASSWORD=deliblogpassword -d postgres:16.3 (username: postgres)