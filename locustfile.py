from locust import HttpLocust, TaskSet, task

class UserBehavior(TaskSet):
    @task(1)
    def view(self):
        self.client.post("/view")

class WebsiteUser(HttpLocust):
    task_set = UserBehavior
    min_wait = 5000
    max_wait = 9000