FROM python:3
RUN pip install requests
COPY trigger.sh /code/trigger.sh
WORKDIR /code
CMD ["python", "-u", "main.py"]