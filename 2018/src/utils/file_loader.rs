pub struct FileLoader {
    pub file_name: String,
}

impl FileLoader {
    pub fn new(file_name: String) -> FileLoader {
        FileLoader { file_name }
    }

    pub fn read(&self) -> String {
        return std::fs::read_to_string(&self.file_name)
            .unwrap_or_else(|_| panic!("Could not read file {}", self.file_name));
    }

    pub fn read_lines(&self) -> Vec<String> {
        let contents = std::fs::read_to_string(&self.file_name)
            .unwrap_or_else(|_| panic!("Could not read file {}", self.file_name));

        contents.split('\n').map(String::from).collect()
    }
}
