#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct {
  bool* floors;
  int state; // 0: static, 1: up, 2: down
  void (*setFloor)(int floor);
} ElevatorPanel;

typedef struct {
  int id;
  int curFloor;
  ElevatorPanel panel;
  void (*go)(int floor);
} Elevator;

typedef struct {
  Elevator* elevators;
  int numElevators;
  void (*run)(int floor);
} ElevatorSystem;

typedef struct {
  int locate;
  int state; // 0: null, 1: up, 2: down
} FloorPanel;

typedef struct {
  int curFloor;
} Passenger;

