# Phase II Testing Plan

## Overview
**Goal**: To capture video footage of robot operation for funding applications

**Participants**: Frank, James, Sean

**Location**: 887 Bay Street

## Objectives
-Robot driving on the streets amidst pedestrians
-Robot "Delivering" pizzas to a client (perhaps Genieve can be our client?) I'll order the pizzas (Bay & Phipps st Pizza Nova) for pickup and we can actually have that for lunch :slightly_smiling_face:
-A mock QR code printed and taped on the box. This will be what our box authentication flow would look like.
-Screen capture (on loom) of our driving interface.
-(Reach goal) convince a restaurant to actually put the food in the robot and film the process.

## Preparation

## Observations
## Test 1: June 21, 2020 @ 1100-1200
-


## Conclusions
### June 21, 2020
Test period: 1 - 1830
Route: Drove around Sean's block
#### Findings:
- Camera image for the Logitech C920 has absolutely no depth perception. I could not determine if there was a curb, a step, or if the ground is flat. This resulted in one incident of the right wheel falling off the curb, and another where I crashed into a set of low stairs.
- Camera vibration is quite intense. When crossing over vents and sidewalk tiles, the entire image shakes and is almost unusable. Falling back to the LiDAR picture really improves driving experience.
- LiDAR create exceptional situational awareness. It's easy to see people passing by and see obstacles and its relative distance to you. From my experience driving the robot, relying on LiDAR for fine navigation and using the camera to inspect the surroundings when needed. I was able to drive alongside other pedestrians and not have any issue identifying their distance based on the LiDAR pings.
- GPS was not working; it is a crucial aspect of the driving experience. Based on the camera image alone, it was really difficult to determine exactly where I was even though I am quite familiar with the area.
- Maintaining lane is not very difficult.
- LiDAR reflections off of foiliage is very good; not so good on objects that are black.
- Speed-level control works, and can get the robot to quite high speeds. It's difficult to judge the speed of the robot by the camera image alone. Having a speed indication based on sensors is necesary!
- The VR2 controller draws variable amounts of current for a given speed level setting suggesting perhaps the speed configuration is a percentage of the top speed, rather than power. The current is NOT constant, and voltage fluctuations on the battery shunt is ~1V when driving and up to 3V when accelerating. **Note:** The fuse (rated 10A) blew out when I hit a curb. This fuse rating is matched with a limitation on the safety relay that we are using.
- There were severe connectivity issues with the robot when it's on the LTE connection; this was not seen when the robot is hooked up to a LAN cable. The error is "tcp i/o timeout" thrown by the Gorilla WebSocket library. This creates an ungraceful disconnect with the signaling server (code 1005 and 1006). Additionally, during this type of disconnect event, the robot does not attempt to register itself with the signaler again, requiring the signaler to be restarted to manually force a robot reconnect.
- The connection state reducer on the frontend should be dispatched when the WebRTC connection is established, rather than when the WebRTC offer has been received. This ensures that the controls are only enabled when the actual data connection(s) are ready.
- A better way of resolving robot registration ID conflicts should be established. The Signaler currently rejects connection from duplicated hosts, so a crashed robot that hasn't deregistered from the Signaler can no longer connect. Should debug why it's not getting deregistered on disconnect/error.
