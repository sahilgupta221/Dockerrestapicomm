# https://hub.docker.com/r/library/python/tags/
FROM python:3-onbuild  

RUN pip install requests

# Coping source in current directory into the image
# python:3-onbuild expects the source in /usr/src/app
COPY . /usr/src/app

# Commands in a list
CMD ["python", "request.py"]

