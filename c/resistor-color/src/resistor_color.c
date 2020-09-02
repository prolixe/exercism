#include "resistor_color.h"

const resistor_band_t expected[] = {
BLACK, BROWN, RED, ORANGE, YELLOW,
GREEN, BLUE, VIOLET, GREY, WHITE
};
int color_code(resistor_band_t r){
	return (int)r;
}

const resistor_band_t* colors(){
   return expected;
}

