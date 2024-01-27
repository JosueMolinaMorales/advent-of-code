use std::{
    fs::File,
    io::{BufReader, Read},
};

pub struct FileLoader {
    pub file_name: String,
}

impl FileLoader {
    pub fn new(file_name: String) -> FileLoader {
        FileLoader { file_name }
    }

    pub fn load_file(&self) -> String {
        let file = File::open(&self.file_name).expect("File not found");
        let mut reader = BufReader::new(file);
        let mut contents = String::new();
        reader
            .read_to_string(&mut contents)
            .expect("Could not read file");
        contents
    }

    pub fn read_lines(&self) -> Vec<String> {
        let contents = std::fs::read_to_string(&self.file_name)
            .expect(format!("Could not read file {}", self.file_name).as_str());

        contents.split("\n").map(|s| String::from(s)).collect()
    }
}
