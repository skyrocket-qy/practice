#include <signal.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "TokenCounterImpl.h"


void my_handler(int s){
    printf("\nCaught signal %d\n", s);
}

std::string getInputString(char* input) {
    size_t length = strcspn(input, "\n");
    input[length] = '\0';
    return std::string(input);
}

int main(int argc, char** argv){
    struct sigaction sigIntHandler;
    sigIntHandler.sa_handler = my_handler;
    sigemptyset(&sigIntHandler.sa_mask);
    sigIntHandler.sa_flags = 0;
    sigaction(SIGINT, &sigIntHandler, NULL);

    TokenCounterImpl* counter = new TokenCounterImpl();

    char input[256];
    while (1) {
        printf("Enter a command: ");
        if (fgets(input, sizeof(input), stdin) == NULL) {
            break;
        }

        input[strcspn(input, "\n")] = '\0';

        if (strlen(input) > 0) {
            if (strncmp(input, "ingest", 6) == 0){
                const std::string str = getInputString(input);
                printf("Actual input: %s\n", str.c_str());
                counter->Ingest(str.substr(6));
            }else if (strncmp(input, "appearance", 10) == 0){
                const std::string str = getInputString(input);
                printf("Actual input: %s\n", str.c_str());
                counter->Appearance(str.substr(10));
            }else{
                printf("Unknown command: %s\n", input);
            }
        }
    }

    printf("clean up the resource\n");
    delete counter;
    return 0;
}

