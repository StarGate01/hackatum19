FROM python:3
COPY requirements.txt /code/requirements.txt
WORKDIR /code
RUN pip install -r requirements.txt
CMD ["python", "app.py"]