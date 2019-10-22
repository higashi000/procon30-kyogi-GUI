#ifndef GAME_HPP_
#define GAME_HPP_

#include <Siv3D.hpp>
#include "button.hpp"
#include "connector.hpp"
#include <vector>
#include <string>
#include "field.hpp"
#include "action.hpp"

namespace p30kG {
class Game {
  private :
    Connector conn;
    Array<Array<Button>> buttons;
    Field field;
    Font font;
    std::vector<bool> isClicked;
    std::vector<Action> actions;
    Rect answerSend;
    Rect fieldUpdate;

  public :
    Game();
    void dispField();
    void setFieldData();
};
}

#endif