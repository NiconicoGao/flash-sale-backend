import http from 'k6/http';
import { Counter } from 'k6/metrics';

const myCounter = new Counter('my_counter');
export const options = {
    stages: [
      { duration: '5s', target: 100 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
      { duration: '10s', target: 200 }, // stay at 100 users for 10 minutes
      { duration: '30s', target: 300 }, // stay at 100 users for 10 minutes
      { duration: '10s', target: 100 }, // ramp-down to 0 users
      { duration: '10s', target: 0 }, // ramp-down to 0 users
    ],
    thresholds: {
      'http_req_duration': ['p(99)<1500'], // 99% of requests must complete below 1.5s
      "my_counter": ['count<100'],
    },
  };
  

export default function () {
    const url = 'http://54.201.152.192/api/order?id=36';
  
    const params = {
      headers: {
        'Content-Type': 'application/text',
      },
      timeout: "5s"
    };
  
    const result = http.get(url,params);
    myCounter.add(result.json('code') === 0)
  }