let x =0;
function spreadsheet(number) {
    for (let i = 10; i == 1; i = i-1) {
        x = x + number + '*' + i + '=' + number*i + '<br>';
      }
      return x;
}
out.innerHTML += spreadsheet(4)

var a=0, b=0, c=0;
for (var i=1; i<=5; i++) {
    b=0;
    a++;
    document.write("<br/>");
    for (var j=1; j<=5; j++){
        b++;
        c = a * b;
        document.write(a + "x" + b + "=" + c + "<br/>");    }
}



let sum =0;
const numbers = [2,14,26,18];
numbers.forEach(fun4);
function fun4(num) 
{
    sum += num;
}
out2.innerHTML += sum;

