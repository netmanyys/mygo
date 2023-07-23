### Ready to go CRUD example


❯ curl http://localhost:8888/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty YAN","price": 49.99}'


❯ curl -X DELETE http://localhost:8888/albums/4
{
    "message": "album deleted"
}%


curl http://localhost:8888/albums/4 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id": "4","title": "Updated Title","artist": "Updated Artist","price": 29.99}'
