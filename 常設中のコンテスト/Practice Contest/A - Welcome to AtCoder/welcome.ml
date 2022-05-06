let () =
	let a = read_int () in
    let string = read_line () in
  	let list = List.map int_of_string (String.split_on_char ' ' string) in
  	let text = read_line() in
  	let ans = a + (List.fold_left (+) 0 list) in
  	Printf.printf "%d %s\n" ans text;
;;

