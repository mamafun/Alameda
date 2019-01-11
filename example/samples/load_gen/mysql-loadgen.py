from locust import HttpLocust, TaskSet, events, task
import  pymysql
import random
import time

class UserBehavior(TaskSet):

    @task
    def get_employees_100(self):
        start_time = time.time()
        try:
            db = pymysql.connect(host="172.31.86.191",port=32169,user="root",passwd="Lxq9XFvRwe",db="employees")
            cursor = db.cursor()
            cursor.execute("SELECT employees.employees.emp_no, employees.employees.first_name, employees.salaries.salary FROM employees.employees, employees.salaries WHERE employees.employees.emp_no = employees.salaries.emp_no and employees.employees.emp_no >= 10000 and employees.employees.emp_no < 10100 order by salary DESC;")
            self.data = cursor.fetchall()
            db.close()
        except Exception as e:
            total_time = int((time.time() - start_time) * 1000)
            events.request_failure.fire(request_type="mysql", name="get_employees_100", response_time=total_time, exception=e)
        else:
            total_time = int((time.time() - start_time) * 1000)
            events.request_success.fire(request_type="mysql", name="get_employees_100", response_time=total_time, response_length=0)

    @task
    def get_employees_5000(self):
        start_time = time.time()
        try:
            db = pymysql.connect(host="172.31.86.191",port=32169,user="root",passwd="Lxq9XFvRwe",db="employees")
            cursor = db.cursor()
            cursor.execute("SELECT employees.employees.emp_no, employees.employees.first_name, employees.salaries.salary FROM employees.employees, employees.salaries WHERE employees.employees.emp_no = employees.salaries.emp_no and employees.employees.emp_no >= 100000 and employees.employees.emp_no < 105000 order by salary DESC;")
            self.data = cursor.fetchall()
            db.close()
        except Exception as e:
            total_time = int((time.time() - start_time) * 1000)
            events.request_failure.fire(request_type="mysql", name="get_employees_5000", response_time=total_time, exception=e)
        else:
            total_time = int((time.time() - start_time) * 1000)
            events.request_success.fire(request_type="mysql", name="get_employees_5000", response_time=total_time, response_length=0)

    @task
    def get_employees_1000(self):
        start_time = time.time()
        try:
            db = pymysql.connect(host="172.31.86.191",port=32169,user="root",passwd="Lxq9XFvRwe",db="employees")
            cursor = db.cursor()
            cursor.execute("SELECT employees.employees.emp_no, employees.employees.first_name, employees.salaries.salary FROM employees.employees, employees.salaries WHERE employees.employees.emp_no = employees.salaries.emp_no and employees.employees.emp_no >= 200000 and employees.employees.emp_no < 201000 order by salary DESC;")
            self.data = cursor.fetchall()
            db.close()
        except Exception as e:
            total_time = int((time.time() - start_time) * 1000)
            events.request_failure.fire(request_type="mysql", name="get_employees_1000", response_time=total_time, exception=e)
        else:
            total_time = int((time.time() - start_time) * 1000)
            events.request_success.fire(request_type="mysql", name="get_employees_1000", response_time=total_time, response_length=0)

class WebsiteUser(HttpLocust):
    host=""
    task_set = UserBehavior
    min_wait = 0
    max_wait = 0
