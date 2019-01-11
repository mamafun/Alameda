from locust import HttpLocust, TaskSet, task
 
class SimpleLocustTest(TaskSet):
 
    @task
    def get_categories(self):
        self.client.get("/applicationPetstore/rest/categories")
 
    @task
    def get_products(self):
        self.client.get("/applicationPetstore/rest/products")

    @task
    def get_items(self):
        self.client.get("/applicationPetstore/rest/items")

    @task
    def get_countries(self):
        self.client.get("/applicationPetstore/rest/countries")

    @task
    def get_customers(self):
        self.client.get("/applicationPetstore/rest/customers")

class LocustTests(HttpLocust):
    task_set = SimpleLocustTest

