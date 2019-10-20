#ifndef CONNECTOR_HPP_
#define CONNECTOR_HPP_

#include <iostream>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string>
#include <vector>
#include <sstream>
#include <unistd.h>
#include "field.hpp"

namespace p30kG {
class Connector {

  public:
    Field getFieldData();
    void sendResult();
};

std::vector<std::string> split(std::string str, char splitChar);

} // namespace p30kG
#endif
