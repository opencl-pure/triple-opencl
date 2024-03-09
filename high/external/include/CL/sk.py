import os


def process_c_headers(directory):
    # Prejdeme všetky súbory v adresári
    for filename in os.listdir(directory):
        if filename.endswith(".h"):  # Kontrola, či je to C hlavičkový súbor
            c_file_path = os.path.join(directory, filename)
            go_file_path = os.path.splitext(c_file_path)[0] + ".go"  # Vytvoríme cestu k výstupnému súboru .go

            with open(c_file_path, 'r') as c_file:
                for line in c_file:
                    words = line.split()
                    if len(words) > 2 and words[0] == "#define":
                        new_line = ""
                        for word in words[1:]:
                            if word != "":
                                if len(new_line) == 0:
                                    new_line = word + " = "
                                    continue
                                new_line += word  # Pridáme slovo
                        if new_line:
                            print(new_line + "\n")  # Zapíšeme nový riadok do .go súboru


# Spustenie skriptu na aktuálnom adresári
process_c_headers('.')
