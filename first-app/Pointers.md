**Pointers**

* accountNumber degiskeninde 300 degerini sakla,

* Dolayisiyla burada iki sey alabilirim; Degiskenin adresi ve degiskenin tuttugu deger

* accountNumberPointer degiskenini *int ile tanitiyorum;

* accountNumberPointer i &accountNumber degiskenine esitlersem accountNumber'a ait adres degerini alabilirim (&ampersand)

* Eger accountNumberPointer'i *accountNumberPointer seklinde ekrana basarsam ekranda 300 degerini goruyorum

* main fonksiyonu icerisinde Account tipinde bir nesne yarattim diger dillerde de oldugu gibi {} araciligi ile Account'u direkt olarak initialize etmis oldum

* Ardindan ChangeAccountNumber ismi methodu cagiriyorum. Bu method icerisinde AccountNumber degiskenini 120'ye esitliyorum ve main methodunda AccountNumber'i print ediyorum. Fakat burada gorulecegi uzere aslinda ChangeAccountNumber methodu ile main fonksiyonu icindeki Account objesi sadece type olarak create ettigimiz Account'un birer ayri kopyasidir.Yani ChangeAccountNumber fonksiyonu main icerisinde ki initialize edilmis olan AccountNumber degiskenini DEGISTIREMEZ

* Printf fonksiyonu icerisinde ki %p ile hexadecimal olarak adresi yazdırabiliyorum

* ChangeAccountNumberWithPointer methoduna &account ile account nesnemi main icerisinden gonderdim, ve ChangeAccountNumberWithPointer isimli methoda parametre olan *Account tipinde yani pointer tipinde degerin gelecegini belirttim. Simdi main fonksiyonu ile ChangeAccountNumberWithPointer methodu aynı degere sahip olduklarindan dolayi AccountNumber'i degistirdigim anda main icerisinde ki degerde degismis olacak.
