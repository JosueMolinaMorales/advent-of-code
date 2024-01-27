pub struct FileLoader {
    pub file_name: String,
}

impl FileLoader {
    pub fn new(file_name: String) -> FileLoader {
        FileLoader { file_name }
    }

    pub fn read_lines(&self) -> Vec<String> {
        let contents = std::fs::read_to_string(&self.file_name)
            .expect(format!("Could not read file {}", self.file_name).as_str());

        contents.split("\n").map(|s| String::from(s)).collect()
    }
}
