from rest_framework import serializers

from book.models import Book


class CreateBookSerializer(serializers.Serializer):
    name = serializers.CharField()


class BookSerializer(serializers.ModelSerializer):
    class Meta:
        model = Book
        fields = "__all__"


class ModifyBookSerializer(CreateBookSerializer):
    id = serializers.IntegerField()


class DeleteBookSerializer(serializers.Serializer):
    id = serializers.IntegerField()
