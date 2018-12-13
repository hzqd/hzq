eat() => print("I want eat.");
colorOfCars() => "car's color is black.";

main(){
  eat();
  print(colorOfCars());
  
  while(true){
    print("1");
    break;
  }

  var number = 1;
  while(5>=number){
    print(number);
    /** number = number + 1; */
    number++;
  }

  for(var i=0;i<10;i++){
    if (i%2==0){
      print("hello $i");
    }
  }

  do{print("do something");}
  while(false);

  //直观感受：分号打得手疼

  var sb = StringBuffer();
  sb
    ..write("hello ")
    ..write("world");
  print(sb.toString());

  var numberReg = RegExp(r'\d+');
  var str = 'Hello 20 and 30';
  for(var match in numberReg.allMatches(str)){
    print(match);
  }

  List<String> list = ["Hello","World"];
  list.add("victory");
  list.addAll(["123","456"]);
  list.forEach((item) => print(item));
  list.removeAt(0);
  print(list.indexOf("456"));
  for(String item in list) {
    print(item);
  }

  Set<int> sets = Set();
  sets.addAll([1,2,3,4]);
  var fn1 = (int item) => print(item);
  sets.forEach(fn1);
  void fn2(int item) => print (item);
  sets.forEach(fn2);

  var map = {
    "name": "qb",
    "age": "99",
    "gender": "mail",
    "other": {
      "a": 1,
      "b": 2,
    }
  };
  print(map);
}