FROM python:3.11.1
ENV PYTHONUNBUFFERED 1
EXPOSE 8000/tcp
COPY python3 /app
RUN useradd -mU -u 50000 -s /bin/bash app && \
  chown -R app /app
WORKDIR /app
RUN pip install -r requirements.txt
USER app
ENTRYPOINT gunicorn proj.wsgi:application --bind 0.0.0.0:8000 --workers 2
