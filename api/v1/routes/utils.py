import os
import logging
from datetime import datetime, timedelta

def get_current_timestamp():
    return datetime.now().timestamp()

def get_file_size_in_bytes(file_path):
    return os.path.getsize(file_path)

def get_timestamp_offset_in_seconds(timestamp, offset):
    return timestamp + offset * 60

def create_logger(name, level=logging.INFO):
    logger = logging.getLogger(name)
    logger.setLevel(level)
    formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    handler = logging.StreamHandler()
    handler.setFormatter(formatter)
    logger.addHandler(handler)
    return logger