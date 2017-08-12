from django.test import Client
from django.test import TestCase
from django.urls import reverse
from rest_framework.test import APIClient


class APITestCase(TestCase):
    client_class = APIClient

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def get_url(self, name):
        return reverse(name)

    def assert_resp(self, resp):
        data = resp.data
        self.assertEqual(data["code"], 200)
        self.assertEqual(data["err"], None)
