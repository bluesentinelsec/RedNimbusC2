# content of test_sample.py

import s3wrapper

bucket_name = s3wrapper.get_s3_test_bucket_name()

def test_put_s3_file():
    ret = s3wrapper.put_s3_file(bucket_name, "test_file.txt", "testing/test_file.txt")
    assert(ret == True)

def test_get_s3_file():
    ret = s3wrapper.get_s3_file(bucket_name, "testing/test_file.txt", "/tmp/test_file.txt")
    assert(ret == True)

def test_list_s3_files():
    ret = s3wrapper.list_s3_files(bucket_name, "testing/")
    assert(ret != "")

def test_remove_s3_file():
    ret = s3wrapper.remove_s3_file(bucket_name, "testing/test_file.txt")
    assert(ret == True)

def test_get_default_region():
    region = s3wrapper.get_default_region()
    assert(region != "")

def test_get_account_id():
    account_id = s3wrapper.get_account_id()
    assert(account_id != "")