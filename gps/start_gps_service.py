#!/usr/bin/env python3

#################################
# Settings
comm_port = '/dev/ttyUSB2'
baudrate = 115200
TIMEOUT = 3
DEBUG = 0

# Libraries
import time
import serial
import os


# Global Variables
ser = serial.Serial()


# Function for printing debug message
def debug_print(message):
    if(DEBUG):
        print(message)
    return

ERROR_CODES = {'501':'Invalid Params',
               '502':'Operation Not Supported',
               '503':'GNSS Subsystem Busy',
               '504':'Session is ongoing',
               '505':'Session not active',
               '506':'Operation Timeout',
               '507':'Function not enabled',
               '508':'Time information error',
               '512':'Validity time is out of range',
               '513':'Internal resource error',
               '514':'GNSS locked',
               '515':'End by E911',
               '516':'Not Fixed -> Need to Acquire Satellites',
               '549':'Unknown Error'}

def decodeError(error):

    error = error.split()[-1]

    if(error in ERROR_CODES.keys()):
        return ERROR_CODES[error]

    return "Error Lookup Failed"


##################################
# GNSS Class
class GNSS:

    def __init__(self, serial_port=comm_port, serial_baudrate=115200):

        ser.port     = serial_port
        ser.baudrate = serial_baudrate
        ser.parity   = serial.PARITY_NONE
        ser.stopbits = serial.STOPBITS_ONE
        ser.bytesize = serial.EIGHTBITS


    # Get Responses from the Modem
    def getResponse(self):

        if (ser.isOpen() == False):
            ser.open()

        response = ser.readline().decode('utf-8')
        if ('ERROR' in response):
            print(decodeError(response))

        return response

    # Send AT Command to the module
    def sendATCommOnce(self, command):

        if (ser.isOpen() == False):
            ser.open()

        self.compose = ""
        self.compose = str(command) + "\r"
        ser.reset_input_buffer()
        ser.write(self.compose.encode())
        ser.readline()
        debug_print(self.compose)

    def configureGNSS(self):
        return self.sendATCommOnce("AT+QGPSCFG=\"gpsnmeatype\",3")

    def turnOnGPS(self):
        return self.sendATCommOnce("AT+QGPS=1")

    def turnOffGPS(self):
        return self.sendATCommOnce("AT+QGPSEND")


    def getPosition(self):
        self.sendATCommOnce("AT+QGPSLOC?")
        response = self.getResponse()

        #response format: <latitude>,<longitude> format: ddmm.mmmm N/S,dddmm.mmmm E/W

        return response

##########################################################
# Main

gps = GNSS()
gps.turnOffGPS()
gps.configureGNSS()
gps.turnOnGPS()
gps.getPosition()

