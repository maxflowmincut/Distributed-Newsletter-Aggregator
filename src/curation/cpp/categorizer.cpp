#include <string>
#include <vector>
#include <map>
#include <set>
#include <algorithm>
#include <cctype>
#include "categorizer.h"

extern "C" {

enum Category {
    SCIENCE,
    TECH,
    ART,
    NONE
};

struct Keyword {
    std::string word;
    double weight;
};

std::vector<Keyword> scienceKeywords = {
    {"biology", 1.5}, {"chemist", 1.5}, {"physicist", 1.5}, {"genome", 2.0},
    {"quantum mechanics", 3.0}, {"molecule", 1.5}, {"organism", 1.5}, {"cell biology", 2.5},
    {"astrophysics", 2.5}, {"neuroscience", 2.5}, {"geology", 1.5}, {"climate change", 2.5},
    {"science", 5.0}, {"climate", 4.0}, {"climate change", 4.0}
};

std::vector<Keyword> techKeywords = {
    {"software", 1.5}, {"hardware", 1.5}, {"computer", 1.0}, {"programming", 1.5},
    {"virtual reality", 2.5}, {"artificial intelligence", 3.0}, {"smartphone", 2.0},
    {"app development", 2.5}, {"cybersecurity", 2.5}, {"augmented reality", 2.5},
    {"machine learning", 2.5}, {"automation", 2.0}, {"twitter", 3.5}, {"facebook", 3.5},
     {"google", 3.5}, {"meta", 3.5}, {"databricks", 3.5}, {"openai", 3.5},  {"robot", 3.5},
     {"tech", 5.0}
};

std::vector<Keyword> artKeywords = {
    {"painting", 2.0}, {"sculpture", 2.0}, {"gallery", 1.5}, {"artist", 1.5},
    {"craftsmanship", 1.5}, {"photography", 2.0}, {"fine arts", 2.5},
    {"cinematography", 2.0}, {"music composition", 2.5}, {"digital art", 2.0},
    {"performing arts", 2.5}, {"literary art", 2.5}, {"art", 5.0}, {"musician", 4.0}
};

std::vector<Keyword> negativeKeywords = {
    {"sale", -2.0}, {"discount", -2.0}, {"advertisement", -2.5}
};

std::string preprocess(const std::string& str) {
    std::string result;
    for (auto& ch : str) {
        if (std::isalnum(ch) || std::isspace(ch)) {
            result += std::tolower(ch);
        }
    }
    return result;
}

double computeScore(const std::string& description, const std::vector<Keyword>& keywords) {
    double score = 0.0;
    std::string processedDescription = preprocess(description);

    for (const auto& keyword : keywords) {
        size_t foundPos = processedDescription.find(keyword.word);
        while (foundPos != std::string::npos) {
            score += keyword.weight;
            if (foundPos > 0 && processedDescription.substr(foundPos - 4, 3) == "not") {
                score -= 2 * keyword.weight;
            }
            foundPos = processedDescription.find(keyword.word, foundPos + 1);
        }
    }

    for (size_t i = 0; i < keywords.size() - 1; i++) {
        for (size_t j = i + 1; j < keywords.size(); j++) {
            size_t pos1 = processedDescription.find(keywords[i].word);
            size_t pos2 = processedDescription.find(keywords[j].word);
            if (pos1 != std::string::npos && pos2 != std::string::npos && std::abs(static_cast<int>(pos1 - pos2)) < 30) {
                score += 1.0;
            }
        }
    }
    return score;
}

int* categorize(const char* description, int* size) {
    std::string desc(description);
    std::set<Category> categoriesSet;  // The original categorize function must be renamed or integrated here

    double negativeScore = computeScore(desc, negativeKeywords);

    if (computeScore(desc, scienceKeywords) + negativeScore > 6.0) {
        categoriesSet.insert(SCIENCE);
    }
    if (computeScore(desc, techKeywords) + negativeScore > 7.0) {
        categoriesSet.insert(TECH);
    }
    if (computeScore(desc, artKeywords) + negativeScore > 6.5) {
        categoriesSet.insert(ART);
    }
    if (categoriesSet.empty()) {
        categoriesSet.insert(NONE);
    }

    *size = categoriesSet.size();
    int* categoriesArray = (int*) malloc(*size * sizeof(int));

    int index = 0;
    for (const Category& cat : categoriesSet) {
        categoriesArray[index++] = static_cast<int>(cat);
    }

    return categoriesArray;
}

void free_categories(int* categories) {
    free(categories);
}

}

extern "C" void my_free(void* ptr) {
    free(ptr);
}