#include <SPI.h>
#include <RH_NRF24.h>
#include <communication.h>

#define NB_DRONE 1

Communication com(9, 10);
bool ok;
uint8_t msgDrone[11];
uint8_t lenMsgDrone = sizeof(msgDrone);
int i = 0;

void setup() {
  Serial.begin(9600);
  ok = com.initRadio();
  delay(1000);
  if(!ok) {
    Serial.println("Init error");
  }
}

void loop() {
  //For each drone write on serial port received message
  for(i = 0; i < NB_DRONE; i++) {
    ok = com.receiveMsg(msgDrone, lenMsgDrone, i + 1);
    if(ok) {
      Serial.print("d:");
      Serial.print(msgDrone[0]);
      Serial.print(";x:");
      Serial.print(msgDrone[1]);
      Serial.print(";y:");
      Serial.print(msgDrone[2]);
      Serial.print(";z:");
      Serial.print(msgDrone[3]);
      Serial.print(";s1:");
      Serial.print(msgDrone[4]);
      Serial.print(";s2:");
      Serial.print(msgDrone[5]);
      Serial.print(";s3:");
      Serial.print(msgDrone[6]);
      Serial.print(";s4:");
      Serial.print(msgDrone[7]);
      Serial.print(";s5:");
      Serial.print(msgDrone[8]);
      Serial.print(";s6:");
      Serial.print(msgDrone[9]);
      Serial.print(";msg:");
      Serial.println(msgDrone[10]);
    }
  }
  delay(400);
} 
