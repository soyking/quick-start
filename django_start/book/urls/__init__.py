from django.conf.urls import url

from book import views

urlpatterns = [
    url(r"^book", views.BookAPI.as_view(), name="book_api"),
]
