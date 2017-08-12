from utils.tests import APITestCase


class BookAPITest(APITestCase):
    def setUp(self):
        super().setUp()
        self.url = self.get_url("book_api")

    def test_create_book(self):
        resp = self.client.post(self.url, data={"name": "book1"})
        self.assert_resp(resp)
        return resp

    def test_get_book(self):
        self.test_create_book()
        resp = self.client.get(self.url)
        self.assert_resp(resp)
        self.assertEqual(len(resp.data["data"]), 1)

    def test_modify_book(self):
        create_resp = self.test_create_book()
        modify_name = "modify"
        data = {"id": create_resp.data["data"]["id"], "name": modify_name}
        resp = self.client.put(self.url, data)
        self.assert_resp(resp)
        self.assertEqual(resp.data["data"]["name"], modify_name)

    def test_delete_book(self):
        create_resp = self.test_create_book()
        data = {"id": create_resp.data["data"]["id"]}
        resp = self.client.delete(self.url, data)
        self.assert_resp(resp)
        resp = self.client.get(self.url)
        self.assert_resp(resp)
        self.assertEqual(len(resp.data["data"]), 0)
