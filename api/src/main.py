import os

from flask import (Flask, jsonify)
from flask_cors import CORS, cross_origin

from app import create_app
from app.archive_search.views import get_archived_replacement_data
from app.web_search.views import get_web_searched_replacement_data
app = create_app('production')


@app.route('/', methods=['POST', 'OPTIONS'])
@cross_origin(origin='*', headers=['Content-Type', 'Authorization'])
def archive_search(request):
    request_json = request.get_json()
    print('request_json')
    print(request_json)
    response = get_archived_replacement_data(
        request_json,
    )
    print('response')
    print(response)
    return jsonify(response)


@app.route('/', methods=['POST', 'OPTIONS'])
@cross_origin(origin='*', headers=['Content-Type', 'Authorization'])
def web_search(request):
    request_json = request.get_json()
    print('request_json')
    print(request_json)
    response = get_web_searched_replacement_data(
        request_json,
    )
    print('response')
    print(response)
    return jsonify(response)


if __name__ == '__main__':
    port = int(app.config.get('PORT', '5000'))
    app.run(debug=True, host='0.0.0.0', port=port)
