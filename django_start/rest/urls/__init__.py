from django.conf.urls import url

from rest import views

urlpatterns = [
    url(r"^rest$", views.RestAPI.as_view(), name="rest_view"),
]
