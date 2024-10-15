# Miniaturized Avionics for Rapid Testing & Handling (MARTHA)

MARTHA is a PCB flight computer that seeks to provide a low cost, easily integratable, and effective solution to collecting flight data from L1/2 certification rockets.
MARTHA features a 9 degree of freedom inertial measurement unit (IMU) as well as an altimeter to detect both altitude and acceleration/roll/orientation changes. MARTHA is the
first PCB I ever made, and I'm proud of it, despite it's faults

## Authors

Kyle Jones & Trevor Miller (for 1.3)

## Version History

* 1.2
    * Final working version that fixed all bugs from previous versions, built in KiCad
    * Works in theory, but manufacturing defect caused dead shorting upon arrival - We debugged for literal weeks and couldn't figure it out
* 1.3
    * Initial version in Altium
    * Switch to flash data storage
    * Rethink entire board layout
    * Refactor dimensions to be 50x50mm
    * Add screw terminals for power/switch
    * Reduce to single sided assembly
    * Move all sensors to SPI

## License

MIT License

## Acknowledgments

Inspired from Phil's Lab KiCad/Altium tutorials - go check out his YouTube channel
