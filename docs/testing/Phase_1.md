# Phase I Testing Plan

## Overview
**Goal**: To prove that the robot can be reliably controlled in an urban setting with people and obstacles

**Participants**: Andrew, Frank, James, Sean

**Location**: Intersection of Rosedale Rd. and Park Rd.

## Objectives
### 1. Observe people's behavior and attitude towards the robot
* How do people react when in proximity of the robot?
* What is the attitude of bystanders towards a moving robot? (hostile? curious? friendly?)
* Do people make way for the robot?
* Robot visibility and reactions while driving around in darker settings

### 2. Control with both camera types (wide angle distorted and undistorted)
* Which camera provides a better sense of speed and depth?
* Is there sufficient peripheral vision while driving with the non-wide angle camera?
* How does the latency feel for either camera?
* Which camera provides better visual information in darker settings?
* Does the camera get blinded by the sun / reflections? Is there lens flare in very bright settings?
* Is the robot controllable at max speed?

### 3. Crossing roads with remote control
* What is the protocol for road-crossing preparation?
* Do the camera provide sufficient angle to look at the road? Or does the robot need to physically turn to "look left and right"?
* Does the robot have sufficient speed to make the crossing in a reasonable amount of time? How long does a crossing on an average street take?
* From "dropping in" on the robot at the controls to understanding the situation and completing a road crossing, how long does the whole process take?


### 4. Obtain more promotional footage of the robot in action
* Film the robot crossing the road
* Film the robot crossing a small road without the leash
* Photos of the cargo box interior (we should try to get a pizza and fit it into the box)
* Photos of the robot with a backdrop of random street people


#### 5. Obtain data on the robot's operation parameters
* How hot does the control box interior get?
* How long does a full charge of the battery last?
* How hot does the interior of the cargo box get?
* How well does the robot handle small curbs and bumps on the road?
* Are there any potential road configurations that will make the robot stuck?
* Acceleration figures and braking distance figures.


## Conclusions
### June 14, 2020
Test period: 1815 - 1830
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
