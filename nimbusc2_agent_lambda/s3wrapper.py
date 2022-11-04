#!/usr/bin/env python3

import logging

import boto3


def get_s3_file(bucket_name, s3_file, dst_file):
    s3_client = boto3.client('s3')
    try:
        s3_client.download_file(bucket_name, Key=s3_file, Filename=dst_file)
    except Exception as e:
        logging.error(e)
        return False

    return True


def put_s3_file(bucket_name, src_file, dst_file):

    s3_client = boto3.client('s3')
    try:
        s3_client.upload_file(src_file, bucket_name, dst_file)
    except Exception as e:
        logging.error(e)
        return False

    return True


def list_s3_files(bucket_name, prefix):

    file_list = []
    s3_client = boto3.client('s3')
    try:
        response = s3_client.list_objects_v2(Bucket=bucket_name, Prefix=prefix)
        for key in response['Contents']:
            file_list.append(key['Key'])
    except Exception as e:
        logging.error(e)
        return ""

    return file_list


def remove_s3_file(bucket_name, s3_file):
    s3_client = boto3.client('s3')
    try:
        s3_client.delete_object(Bucket=bucket_name, Key=s3_file)
    except Exception as e:
        logging.error(e)
        return False

    return True


def get_s3_bucket_name():
    region = get_default_region()
    account_id = get_account_id()
    bucket_name = f"red-nimbus-c2-{region}-{account_id}"
    return bucket_name


def get_s3_test_bucket_name():
    region = get_default_region()
    account_id = get_account_id()
    bucket_name = f"red-nimbus-c2-testing-{region}-{account_id}"
    return bucket_name


def get_default_region():
    region = ""
    try:
        our_session = boto3.session.Session()
        region = our_session.region_name
    except Exception as e:
        logging.error(e)
        return ""

    return region


def get_account_id():
    account_id = ""
    try:
        account_id = boto3.client('sts').get_caller_identity().get('Account')
    except Exception as e:
        logging.error(e)
        return ""

    return account_id
