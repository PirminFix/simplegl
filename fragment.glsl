#version 130

in vec2 Texcoord;

out vec4 outColor;

uniform sampler2D texCat;

uniform float time;

void main() {
	vec2 coord;
	float x, y;
	x = Texcoord[0];
	y = Texcoord[1];

        if (y > 0.5) {
                x = x - 0.01 * sin((y + time/10) * 100);
                y = max(1.0 - abs(y - 0.5), 0.0) - 0.5;
                coord = vec2(x, y);
                outColor = texture(texCat, coord) * vec4(0.7, 0.7, 0.7, 1.0);
        } else {
                coord = vec2(x, y);
                outColor = texture(texCat, coord);
        }
}
