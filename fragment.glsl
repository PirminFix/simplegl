#version 130

in vec2 Texcoord;

out vec4 outColor;

uniform sampler2D texCat;
uniform sampler2D texDog;

uniform float time;

void main() {

	vec4 colCat = texture(texCat, Texcoord);
	vec4 colDog = texture(texDog, Texcoord);
	outColor = mix(colDog, colCat, sin(time));
}
