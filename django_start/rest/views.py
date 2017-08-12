from rest.models import Book
from rest.serializers import CreateBookSerializer, BookSerializer, ModifyBookSerializer, DeleteBookSerializer
from utils.api import validate, APIView, APIError


class RestAPI(APIView):
    @validate(CreateBookSerializer)
    def post(self, request):
        data = request.json_data
        book = Book.objects.create(name=data["name"])
        return self.response(BookSerializer(book))

    def get(self, request):
        books = Book.objects.all()
        return self.response(BookSerializer(books, many=True))

    @validate(ModifyBookSerializer)
    def put(self, request):
        data = request.json_data
        try:
            book = Book.objects.get(id=data["id"])
            book.name = data["name"]
            book.save()
            return self.response(BookSerializer(book))
        except Book.DoesNotExist:
            raise APIError(code=404, err="book not exist")

    @validate(DeleteBookSerializer)
    def delete(self, request):
        data = request.json_data
        Book.objects.filter(id=data["id"]).delete()
        return self.response_raw()
