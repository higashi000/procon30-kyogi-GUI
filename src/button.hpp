#ifndef BUTTON_HPP_
#define BUTTON_HPP_

#include <Siv3D.hpp>

namespace p30kG {
class Button {
  public :
    Rect rect;
  private :
    int nowColor;
    int isArea;
    String point;
    Font font;
    Vec2 pos;

  public:
    Button();
    Button(int color, int area, int x, int y, String tilePoint);
    void updateColor();
    Button& draw();
};
} // namespace p30kG

#endif