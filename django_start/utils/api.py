from rest_framework.parsers import JSONParser
from rest_framework.response import Response
from rest_framework.views import APIView as RestAPIView


class APIError(Exception):
    def __init__(self, err=None, code=400):
        self.code = code
        self.err = err


class APIView(RestAPIView):
    parser_classes = (JSONParser,)

    def response(self, serializer, code=200):
        return self.response_raw(data=serializer.data, code=code)

    def response_raw(self, data=None, err=None, code=200):
        return Response({"code": code, "err": err, "data": data})

    def dispatch(self, request, *args, **kwargs):
        try:
            return super(APIView, self).dispatch(request, *args, **kwargs)
        except APIError as e:
            response = self.response_raw(err=e.err, code=e.code)
            return self.finalize_response(request, response, *args, **kwargs)


def validate(serializer, many=False):
    def _validate(handler):
        def handle(*args, **kwargs):
            request = args[1]
            result = serializer(data=request.data, many=many)
            if result.is_valid():
                request.json_data = result.data
                return handler(*args, **kwargs)
            else:
                raise APIError(code=400, err=result.errors)

        return handle

    return _validate
