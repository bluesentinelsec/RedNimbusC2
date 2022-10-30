import json
import logging
#import boto3


def handler(event, context):
    # setup logging
    logging.basicConfig(level=logging.INFO,
                        format='%(levelname)s (%(filename)s:%(lineno)s) %(message)s')
    
    # print new requests to console
    logging.info("received new request:")
    print('request: {}'.format(json.dumps(event)))

    # route request to correct function
    
    # return response to agent
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'text/plain'
        },
        'body': 'Hello, CDK! You have hit {}\n'.format(event['path'])
    }


def route_request(path: str):
    pass

def handle_get_task():
    pass


def delete_queued_task():
    pass


def handle_post_task_output():
    pass


def register_agent():
    pass


def remove_agent():
    pass


def is_agent_registered():
    pass
