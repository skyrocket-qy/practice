#include <signal.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

volatile sig_atomic_t ctrlCPressed = 0;

void my_handler(int s){
    printf("\nCaught signal %d\n", s);
    ctrlCPressed = 1;
}

int main(int argc, char** argv){
    struct sigaction sigIntHandler;

    sigIntHandler.sa_handler = my_handler;
    sigemptyset(&sigIntHandler.sa_mask);
    sigIntHandler.sa_flags = 0;

    sigaction(SIGINT, &sigIntHandler, NULL);

    char input[256];
    while (1) {
        printf("Enter a command: ");
        if (fgets(input, sizeof(input), stdin) == NULL) {
            break;
        }

        input[strcspn(input, "\n")] = '\0';

        if (ctrlCPressed) {
            printf("Ctrl+C pressed. Program will exit if 'exit' is entered.\n");
            ctrlCPressed = 0;
        } else if (strlen(input) > 0) {
            printf("Unknown command: %s\n", input);
        }
    }

    return 0;
}
