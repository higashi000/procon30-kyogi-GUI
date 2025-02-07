// MIT-License https://github.com/higashi000/procon30-kyogi
import std.stdio, std.conv, std.string, std.array;
import Kanan.field, Kanan.dispField, Kanan.montecarlo, Kanan.beamSearch, Kanan.connector, Kanan.sendData;
import core.thread;

// コマンドライン引数にMariのip,port,試合の最大ターン数をつけて
void main(string[] args)
{
  auto connector = new KananConnector(args[1], args[2].to!ushort);
  uint turn = 1;
  uint maxTurn = args[3].to!uint;


  while (turn <= maxTurn) {
    Thread.sleep(dur!("seconds")(10));
    connector.getFieldData();
    auto field = connector.parseFieldData();
    writeln(turn);
    disp(field);
    writeln;
    field.calcTilePoint;
    field.calcMyAreaPoint;
    field.calcRivalAreaPoint;

    writeln(field.myTilePoint + field.myAreaPoint);
    writeln(field.rivalTilePoint + field.rivalAreaPoint);

    auto montecarlo = new MontecarloTreeSearch(field, turn, maxTurn, to!int(args[4]), 20);
    auto answer = montecarlo.playGame();
    writeln(answer);
    connector.sendResult(answer);


    turn++;
  }
  auto field = connector.parseFieldData();
  disp(field);
  writeln;
  field.calcTilePoint;
  field.calcMyAreaPoint;
  field.calcRivalAreaPoint;
  writeln(field.myTilePoint + field.myAreaPoint);
  writeln(field.rivalTilePoint + field.rivalAreaPoint);
}
