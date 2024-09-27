#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

// Task structure to hold the function pointer and argument
typedef struct {
    void (*function)(void* arg);  // Function with a generic signature
    void* arg;                    // Argument to be passed to the function
} Task;

// Task queue and thread pool
#define MAX_THREADS 4
#define MAX_TASKS 10

Task taskQueue[MAX_TASKS];
int taskCount = 0;
pthread_mutex_t taskMutex;
pthread_cond_t taskCond;

void* threadWorker(void* arg) {
    while (1) {
        Task task;

        pthread_mutex_lock(&taskMutex);
        // Wait until there are tasks in the queue
        while (taskCount == 0) {
            pthread_cond_wait(&taskCond, &taskMutex);
        }

        // Get the task from the queue
        task = taskQueue[--taskCount];
        pthread_mutex_unlock(&taskMutex);

        // Execute the task
        task.function(task.arg);  // Call the task function with the argument
    }
    return NULL;
}

void initThreadPool(pthread_t* threads) {
    pthread_mutex_init(&taskMutex, NULL);
    pthread_cond_init(&taskCond, NULL);

    for (int i = 0; i < MAX_THREADS; ++i) {
        pthread_create(&threads[i], NULL, threadWorker, NULL);
    }
}

// Add a task to the task queue
void addTask(void (*function)(void*), void* arg) {
    pthread_mutex_lock(&taskMutex);
    if (taskCount < MAX_TASKS) {
        Task task;
        task.function = function;
        task.arg = arg;
        taskQueue[taskCount++] = task;

        pthread_cond_signal(&taskCond);  // Signal a thread to wake up and handle the task
    }
    pthread_mutex_unlock(&taskMutex);
}

// Wrapper for a task that takes an int argument
void intTaskWrapper(void* arg) {
    int taskId = *(int*)arg;
    printf("Int Task %d is being processed\n", taskId);
    sleep(1);
}

// Wrapper for a task that takes a double argument
void doubleTaskWrapper(void* arg) {
    double taskValue = *(double*)arg;
    printf("Double Task with value %.2f is being processed\n", taskValue);
    sleep(1);
}

int main() {
    pthread_t threads[MAX_THREADS];
    initThreadPool(threads);

    // Add tasks with different argument types
    for (int i = 0; i < 5; ++i) {
        int* intTaskId = malloc(sizeof(int));  // Allocate memory for int task ID
        *intTaskId = i + 1;
        addTask(intTaskWrapper, intTaskId);    // Use the intTaskWrapper for int tasks
    }

    for (int i = 0; i < 5; ++i) {
        double* doubleTaskValue = malloc(sizeof(double));  // Allocate memory for double value
        *doubleTaskValue = (i + 1) * 1.5;
        addTask(doubleTaskWrapper, doubleTaskValue);       // Use the doubleTaskWrapper for double tasks
    }

    // Allow time for threads to process tasks
    sleep(10);

    return 0;
}
// https://learnku.com/articles/41728