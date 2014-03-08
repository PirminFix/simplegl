#version 130

in vec2 Texcoord;

out vec4 outColor;

uniform sampler2D texCat;


void main() {
	vec2 coord;
	float x, y;
	x = Texcoord[0];
	y = Texcoord[1];
	coord = vec2(x, max(1 - abs(y - 0.5), 0) - 0.5);
	outColor = texture(texCat, coord);
}
