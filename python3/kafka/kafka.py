import kafka as dpkp  # type: ignore
import msgpack

# https://kafka-python.readthedocs.io/en/master/

def new_dpkp_producer() -> dpkp.KafkaProducer:  # type: ignore # pylint: disable=E1101
  """
  Return the kafka producer

  Example:

  from kafka import new_dpkp_producer
  p = new_dpkp_producer()
  p.send(topic, b'some_message_bytes')
  """
  return dpkp.KafkaProducer(  # type: ignore # pylint: disable=E1101
    bootstap_servers="broker:9092",
    client_id="abnml_producer",
    retries=5,
    security_protocol="PLAINTEXT",
  )


def new_dpkp_consumer(topic: str) -> dpkp.KafkaConsumer:  # type: ignore # pylint: disable=E1101
  """
  Return the kafka consumer
  
  Example:

  from kafka import new_dpkp_consumer
  c = new_dpkp_consumer(topic)
  for msg in consumer:
    assert isinstance(msg.value, dict)
  """
  result = dpkp.KafkaConsumer(
    bootstap_servers="broker:9092",
    client_id="abnml_consumer",
    enable_auto_commit=True,
    group_id="abnml_consumer_group",
    value_deserializer=msgpack.loads,
  )
  result.subscribe(topics=[topic])
  return result

